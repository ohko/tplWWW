<template>
    <DefaultLayout>
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h4 class="modal-title"><i class="fa fa-success"></i>成功</h4>
                </div>
                <div class="modal-body">
                    <p><code>{[{ . }]}</code></p>
                </div>
                <div class="modal-footer">
                    <a href="/">首页</a> |
                    <a href="/admin">Admin</a>
                    <a href="javascript:history.back()" class="btn btn-default pull-left">返回</a>
                </div>
            </div>
        </div>
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
        mounted() { },
        destroyed() { },
        methods: {},
        filters: {},
        watch: {},
        computed: {},
    }
</script>