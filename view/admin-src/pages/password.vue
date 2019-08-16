<template>
    <DefaultLayout>
        <section class="content-header">
            <div class="row">
                <div class="col-xs-12">
                    <!-- general form elements -->
                    <div class="box box-primary">
                        <div class="box-header with-border">
                            <h3 class="box-title">修改密码</h3>
                        </div>
                        <!-- /.box-header -->
                        <!-- form start -->
                        <form role="form" @submit.prevent="chk">
                            <div class="box-body">
                                <div class="form-group">
                                    <label>新密码</label>
                                    <input type="password" class="form-control" placeholder="新密码" v-model="form.Pass">
                                </div>
                                <div class="form-group">
                                    <label>确认密码</label>
                                    <input type="password" class="form-control" placeholder="确认密码" v-model="form.Pass2">
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
                    Pass: "",
                    Pass2: "",
                }
            }
        },
        props: {},
        created() { },
        mounted() { },
        destroyed() { },
        methods: {
            chk() {
                if (this.form.Pass.length == 0 || this.form.Pass != this.form.Pass2) {
                    alert("两次输入的密码不一致！")
                    return false
                }
                return true
            }
        },
        filters: {},
        watch: {},
        computed: {},
    }
</script>