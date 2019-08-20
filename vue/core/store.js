// 状态管理

module.exports = new Vuex.Store({
   state: {
      clickCount: parseInt(localStorage.getItem("clickCount")) || 0
   },
   mutations: {
      incrementClickCount(state, o) {
         state.clickCount += parseInt(o.x)
         localStorage.setItem("clickCount", state.clickCount)
      }
   }
})