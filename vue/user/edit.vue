<template>
    <DefaultLayout>
        <section class="content-header">
            <div class="row">
                <div class="col-xs-12">
                    <!-- general form elements -->
                    <div class="box box-primary">
                        <div class="box-header with-border">
                            <h3 class="box-title">修改用户</h3>
                        </div>
                        <!-- /.box-header -->
                        <!-- form start -->
                        <form role="form" @submit.prevent="submit">
                            <div class="box-body">
                                <div class="form-group">
                                    <label>帐号</label>
                                    <input type="text" class="form-control" placeholder="输入用户名" v-model="form.User"
                                        readonly>
                                </div>
                                <div class="form-group">
                                    <label>密码</label>
                                    <input type="password" class="form-control" placeholder="密码" v-model="form.Pass">
                                </div>
                                <div class="form-group">
                                    <label>邮箱</label>
                                    <input type="email" class="form-control" placeholder="输入邮箱地址" v-model="form.Email">
                                </div>
                            </div>
                            <!-- /.box-body -->

                            <div class="box-footer">
                                <button type="submit" class="btn btn-primary">提交</button>
                            </div>
                        </form>
                    </div>
                    <!-- /.box -->

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
                form: {
                    ID: 0,
                    User: "",
                    Pass: "",
                    Email: "",
                }
            }
        },
        props: {},
        created() { },
        mounted() {
            this.$getJSON('/admin_user/detail', { ID: this.$route.params.id }, x => {
                if (x.no != 0) return toastr.error(x.msg)
                this.form = x.data
            })
        },
        destroyed() { },
        methods: {
            submit() {
                this.$post('/admin_user/edit', this.form, x => {
                    if (x.no != 0) return toastr.error(x.msg)
                    this.$router.push('/admin/user/list')
                })
            }
        },
        filters: {},
        watch: {},
        computed: {},
    }
</script>