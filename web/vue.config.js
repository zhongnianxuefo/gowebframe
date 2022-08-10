const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  outputDir: '../server/web',
  devServer: {
    port: 8065,   //指定端口号以侦听
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8060/api', // 请求接口地址
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      },
    }
  }
})
