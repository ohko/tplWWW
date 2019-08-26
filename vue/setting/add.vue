<template>
    <DefaultLayout>
        <section class="content-header">
            <div class="row">
                <div class="col-xs-12">
                    <!-- general form elements -->
                    <div class="box box-primary">
                        <div class="box-header with-border">
                            <h3 class="box-title">添加配置</h3>
                        </div>
                        <!-- /.box-header -->
                        <!-- form start -->
                        <form role="form" @submit.prevent="submit">
                            <div class="box-body">
                                <div class="form-group">
                                    <label>关键词</label>
                                    <input type="text" class="form-control" placeholder="输入关键词" v-model="form.Key">
                                </div>
                                <div class="form-group">
                                    <label>说明</label>
                                    <input type="text" class="form-control" placeholder="输入说明" v-model="form.Desc">
                                </div>
                                <div class="form-group">
                                    <label>数字型</label>
                                    <input type="number" class="form-control" placeholder="输入数字型" v-model="form.Int">
                                </div>
                                <div class="form-group">
                                    <label>逻辑型</label>
                                    <input type="checkbox" v-model="form.Bool">
                                </div>
                                <div class="form-group">
                                    <label>字符串</label>
                                    <input type="text" class="form-control" placeholder="输入字符串" v-model="form.String">
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
                    Key: "",
                    Desc: "",
                    Int: 0,
                    String: "",
                    Bool: false,
                }
            }
        },
        props: {},
        created() { },
        mounted() { },
        destroyed() { },
        methods: {
            submit() {
                this.$post("/admin_setting/add", this.form, x => {
                    if (x.no != 0) return alert(x.data);
                    this.$router.push('/admin/setting/list');
                })
            }
        },
        filters: {},
        watch: {},
        computed: {},
    }
</script>