<template>
    <DefaultLayout>
        <section class="content-header">
            <router-link to="/admin/user/add" class="btn btn-primary">添加</router-link>
            <div class="box">
                <div class="box-header">
                    <h3 class="box-title">Users</h3>
                </div>
                <div class="box-body">
                    <table class="table table-hover" id="list">
                        <thead>
                            <tr>
                                <th>User</th>
                                <th>Email</th>
                                <th>-</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="user in list">
                                <td>{{user.User}}</td>
                                <td>{{user.Email}}</td>
                                <td>
                                    <router-link :to="'/admin/user/edit/'+user.User" class="btn btn-primary">编辑
                                    </router-link>
                                    <div @click="del(user.User)" class="btn btn-danger">删除</div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </section>
    </DefaultLayout>
</template>

<script>
    import DefaultLayout from "../layout/default.vue";

    export default {
        components: {
            DefaultLayout
        },
        data() {
            return {
                list: []
            }
        },
        props: {},
        created() { },
        mounted() {
            this.refresh()
        },
        destroyed() { },
        methods: {
            refresh() {
                this.$getJSON("/admin_user/list", null, x => {
                    if (x.no != 0) return alert(x.msg)
                    this.list = x.data

                    setTimeout(_ => {
                        window.xx = $('#list').DataTable({
                            'paging': true,
                            'lengthChange': true,
                            'searching': true,
                            'ordering': true,
                            'info': true,
                            'autoWidth': true
                        })
                    })
                })
            },
            del(user) {
                if (!confirm('确定删除吗？')) return;

                this.$getJSON('/admin_user/delete', { User: user }, x => {
                    if (x.no != 0) return alert(x.msg)
                    $("#list").DataTable().destroy()
                    this.refresh()
                })
            }
        },
        filters: {},
        watch: {},
        computed: {},
    }
</script>