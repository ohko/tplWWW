<template>

</template>

<script>
  import { getJSON } from "../core/request.js"

  export default {
    components: {},
    data() { return {} },
    props: {},
    created() {
      this.$getJSON("/admin/logout", null, x => {
        if (x.no != 0) return alert(x.data);
        this.$router.push('/admin/login')
      });
      return false;
    },
    mounted() { },
    destroyed() { },
    methods: {},
    filters: {},
    watch: {},
    computed: {},
  }
</script>