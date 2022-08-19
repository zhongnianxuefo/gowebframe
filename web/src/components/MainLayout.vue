<template>
  <a-layout>
    <a-layout-sider v-model:collapsed="collapsed" ></a-layout-sider>
    <a-layout-sider
        v-model:collapsed="collapsed"
        :trigger="null"
        collapsible
        :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0, top: 0, bottom: 0 }"
    >
      <div class="logo" />
      <a-menu v-model:selectedKeys="selectedKeys" v-model:openKeys="openKeys" theme="dark" mode="inline">
        <side-menu
            v-for="menu in allMenus"
            :key="menu.title"
            :menu-title="menu.title"
            :menu-key="menu.menuKey"
            :menu-icon="menu.icon"
            :hasSubmenu="menu.hasSubmenu"
            @click="onMenuClick(menu)"
        >
          <side-menu
              v-for="m in menu.subMenus"
              :key="m.title"
              :menu-title="m.title"
              :menu-key="m.menuKey"
              :menu-icon="m.icon"
              @click="onMenuClick(m)"
          />
        </side-menu>
      </a-menu>
    </a-layout-sider>

    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <menu-unfold-outlined
            v-if="collapsed"
            class="trigger"
            @click="setCollapsed(false)"
        />
        <menu-fold-outlined
            v-else
            class="trigger"
            @click="setCollapsed(true)"
        />
        <div  class="header-right">
        </div>
      </a-layout-header>
      <a-layout-content
          :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '880px' }"
      >
        <slot></slot>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { MenuUnfoldOutlined, MenuFoldOutlined} from '@ant-design/icons-vue';
import { ref, defineProps } from 'vue';
import router from "../router"
import SideMenu from "./SideMenu.vue"
import {useRoute} from "vue-router";
import { useStore } from 'vuex'
import request from "@/request";

// const route = useRoute()
const store = useStore()

const allMenus = ref([
  {title:"首页", icon:"home-outlined", page:"home", menuKey:"home", hasSubmenu:false},
]);
allMenus.value = getAllMenuPara()
const collapsed = ref<boolean>(store.getters["setting/getCollapsed"]);
console.log(store.getters["setting/getCollapsed"])

const selectedKey = getSelectedKey();
const selectedKeys = ref<string[]>([selectedKey]);
const openKey = getOpenMenuKey(allMenus.value,selectedKey,collapsed.value);
const openKeys = ref<string[]>([openKey]);


// console.log(global)
// window.localStorage.setItem("menus", menus)
// Vue.ls.set ("menus", menus)
// console.log("menus",menus);
// allMenus.value = doMenuPara(menus)
// console.log("allMenus", menus, allMenus.value);



const setCollapsed= (c:boolean) => {
  store.commit("setting/setCollapsed", c)
  collapsed.value = c;

}

const onMenuClick= (menu:any) => {
  console.log(menu);
  let page = menu["page"];
  if (page!=undefined){
    let module = menu["module"];
    if (module==undefined){
      router.push({name:page});
    }else{
      router.push({name:page, params:{module:module}});
    }
  }
};





function getAllMenuPara():any{
  const store = useStore()
  let menus = store.getters["account/getAllMenus"]
  return doMenuPara(menus);
}

function doMenuPara(menus :any):any{
  let menuParas = [];
  for (let i=0;i<menus.length;i++){
    let m = menus[i];
    let title = m["title"];
    let page = m["page"] ;
    let module = m["module"];
    let subMenus = m["subMenus"];
    let menuKey = ""
    if (page==undefined){
      menuKey = title;
    }else{
      menuKey =  module==undefined ?  page : page+"-"+module;
    }
    m["menuKey"] = menuKey;
    m["hasSubmenu"] = false;
    if (subMenus!=undefined){
      m["hasSubmenu"] = true;
      m["subMenus"] = doMenuPara(m["subMenus"])
      m["icon"] =  m["icon"]!=undefined ?  m["icon"]: "bars-outlined";
    }
    menuParas.push(m)
  }
  return menuParas;
}
function getSelectedKey():string{
  const route = useRoute()
  let selectedKey = route.name as string;

  const module = ref( route.params["module"]);
  if (module.value != undefined){
    selectedKey = selectedKey + '-' + module.value;
  }
  console.log("SelectedKey",selectedKey)
  return selectedKey;
}
function getOpenMenuKey(allMenus:any,selectedKey:string,collapsed:boolean):string {
  let key=""
  if (collapsed==false){
    for (let i=0;i<allMenus.length;i++) {
      let m = allMenus[i];
      let k = m["menuKey"];
      let subMenus =m["subMenus"] ;
      if (subMenus != undefined){
        if (checkSubMenusKey(m["subMenus"] , selectedKey)) {
          key =  k
        }
      }

    }
  }
  console.log("OpenMenuKey",key)
  return key;
}
// const selectedKey =<string> route.name + '-' + route.params["module"];
// console.log(selectedKey,module)
// let col = route.query["collapsed"]
// if (col=="true"){
//   collapsed.value=true
// }else{
//   let openKey = getOpenMenuKey("", allMenus.value, selectedKey)
//   console.log("openKey", openKey);
//   openKeys.value.push (openKey)
// }
// console.log(collapsed)



function checkSubMenusKey(subMenus:any, selectedKey:string):boolean{

  for (let i=0;i<subMenus.length;i++) {
    let m = subMenus[i];
    let k = m["menuKey"];
    if (k == selectedKey) {
      return true
    }
  }
  return false;
}

</script>
<style>

.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger:hover {
  color: #1890ff;
}

.logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.3);
  margin: 16px;
}

.site-layout .site-layout-background {
  background: #fff;
}

.header-right {
  float: right;
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  color: inherit;
}
</style>
