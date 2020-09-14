<template>
   <div id="app">{{count}}
      <button @click="inc">click</button>
   </div>
</template>

<script>
   export default {
      components: {},
      data() {
         return {
            count: 0
         }
      },
      props: {},
      created() { },
      mounted() { },
      destroyed() { },
      methods: {
         inc() {
            this.count++
         }
      },
      filters: {},
      watch: {},
      computed: {},
   }
</script>