module.exports = {
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
        args[0].title = 'Todo'
        return args
      })
  },
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      '^/api': {
       target: 'http://localhost:8082',
        changeOrigin: true,
        logLevel: 'debug',

      },
    },
  },
}
