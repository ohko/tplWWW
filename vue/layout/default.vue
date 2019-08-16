<template>
   <div class="wrapper">

      <header class="main-header">
         <a class="logo" data-toggle="push-menu" role="button">
            <span class="logo-mini">ADM</span>
            <span class="logo-lg">Admin</span>
         </a>
      </header>

      <aside class="main-sidebar">
         <section class="sidebar">
            <ul class="sidebar-menu" data-widget="tree">
               <li class="header">MAIN NAVIGATION</li>
               <template v-for="(m,k) in menus">
                  <li class="treeview" v-if="m.Child" :menuID="k">
                     <a>
                        <i :class="'fa '+m.Class"></i>
                        <span>{{m.Text}}</span>
                        <span class="pull-right-container">
                           <i class="fa fa-angle-left pull-right"></i>
                        </span>
                     </a>
                     <ul class="treeview-menu">
                        <li v-for="c in m.Child">
                           <a :href="c.Href">
                              <i :class="'fa '+c.Class"></i> {{c.Text}}</a>
                        </li>
                     </ul>
                  </li>
                  <li v-if="!m.Child">
                     <a :href="m.Href">
                        <i :class="'fa '+m.Class"></i>
                        <span>{{m.Text}}</span>
                     </a>
                  </li>
               </template>
            </ul>
         </section>
      </aside>

      <div class="content-wrapper">
         <slot></slot>
      </div>

      <footer class="main-footer">
         <div class="pull-right hidden-xs">
            <b>Version</b> 1.0.0
         </div>
         <strong>Copyright &copy; 2019.</strong> All rights reserved.
      </footer>

      <div class="control-sidebar-bg"></div>
   </div>
</template>

<script>
   export default {
      components: {},
      data() {
         return {
            menus: []
         }
      },
      props: {},
      beforeCreate() {
         setTimeout(_ => {
            $('.main-sidebar').tree()
            $(".content-wrapper").height(window.innerHeight - 51)
            $(".sidebar a").each(function (i, a) {
               if (a.href.split("?")[0] == location.href.split("?")[0]) {
                  $(a).parents("li").addClass("active menu-open");
               }
            });
         })
      },
      created() {
         this.$getJSON('/admin/get_adm_menu', null, x => {
            if (x.no != 0) return alert(x.data);
            this.menus = x.data;
         });
      },
      mounted() { },
      destroyed() { },
      methods: {
         logout() {
            this.$getJSON('/admin/logout', null, x => {
               if (x.no != 0) return alert(x.data);
               this.$root.go('/admin/login')
            });
         }
      },
      filters: {},
      watch: {},
      computed: {},
   }
</script>