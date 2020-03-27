<template>
    <DefaultLayout>
        <section class="content-header">

            <div class="box">
                <div class="box-header">
                    <h3 class="box-title">Setting</h3>
                    <div class="box-tools">
                        <router-link to="/admin/setting/add" class="btn btn-sm btn-primary">添加</router-link>
                    </div>
                </div>
                <div class="box-body">
                    <table class="table table-hover" id="list">
                        <thead>
                            <tr>
                                <th>Key</th>
                                <th>Desc</th>
                                <th>Int</th>
                                <th>String</th>
                                <th>Bool</th>
                                <th>-</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="setting in list">
                                <td>{{setting.Key}}</td>
                                <td>{{setting.Desc}}</td>
                                <td>{{setting.Int}}</td>
                                <td>{{setting.String}}</td>
                                <td>{{setting.Bool}}</td>
                                <td>
                                    <router-link :to="'/admin/setting/edit/'+setting.Key" class="btn btn-primary">编辑
                                    </router-link>
                                    <div @click="del(setting.Key)" class="btn btn-danger">删除</div>
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
        data() { return { list: [] } },
        props: {},
        created() { },
        mounted() {
            this.refresh()
        },
        destroyed() { },
        methods: {
            refresh() {
                this.$getJSON("/admin_setting/list", null, x => {
                    if (x.no != 0) return alert(x.msg)
                    this.list = x.data
                })
            },
            del(key) {
                if (!confirm('确定删除吗？')) return;

                this.$getJSON('/admin_setting/delete', { Key: key }, x => {
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