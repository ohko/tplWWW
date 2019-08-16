<template>
    <DefaultLayout>
        <section class="content-header">
            <a href="/admin/user/add" class="btn btn-primary">添加</a>
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
                            {[{ range .us}]}
                            <tr>
                                <td>{[{.User}]}</td>
                                <td>{[{.Email}]}</td>
                                <td>
                                    <!-- <a href="/admin/user/edit?User={[{.User}]}" class="btn btn-primary">编辑</a> -->
                                    <!-- <a href="/admin/user/delete?User={[{.User}]}" onclick="return confirm('确定删除吗？')"
                                        class="btn btn-danger">删除</a> -->
                                </td>
                            </tr>
                            {[{ end }]}
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
        data() { return {} },
        props: {},
        created() { },
        mounted() {
            $('#list').DataTable({
                'paging': true,
                'lengthChange': true,
                'searching': true,
                'ordering': true,
                'info': true,
                'autoWidth': true
            })
        },
        destroyed() { },
        methods: {},
        filters: {},
        watch: {},
        computed: {},
    }
</script>