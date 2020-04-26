<template>
    <DefaultLayout>
        <section class="content-header">
            <div class="row">
                <div class="col-md-12">
                    <div class="box">
                        <div class="box-header">
                            <h3 class="box-title">Users</h3>
                            <div class="box-tools">
                                <router-link to="/admin/user/add" class="btn btn-sm btn-primary">添加</router-link>
                            </div>
                        </div>
                        <div class="box-body">
                            <div id="myGrid" class="ag-theme-balham" style="width: 100%;height:600px;"></div>
                        </div>
                    </div>
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
                agGrid: null,
            }
        },
        props: {},
        created() { },
        mounted() {
            var gridOptions = {
                columnDefs: [
                    { field: 'ID', checkboxSelection: true },
                    { field: 'User' },
                    { field: 'Email' },
                ],

                defaultColDef: {
                    flex: 1,
                    // sortable: true,
                    // filter: true,
                    resizable: true,
                },
                // floatingFilter: true,
                rowModelType: 'serverSide',
                animateRows: true,
                cacheBlockSize: 100,
                maxBlocksInCache: 10,
                enableRangeSelection: true,
                rowSelection: 'multiple',
                allowContextMenuWithControlKey: true,
                getContextMenuItems: params => {
                    return [
                        'chartRange',
                        'separator',
                        {
                            name: '删除',
                            action: this.del,
                        },
                    ]
                },
                statusBar: {
                    statusPanels: [
                        // { statusPanel: 'agTotalAndFilteredRowCountComponent', align: 'left' },
                        // { statusPanel: 'agTotalRowCountComponent', align: 'center' },
                        // { statusPanel: 'agFilteredRowCountComponent' },
                        { statusPanel: 'agSelectedRowCountComponent' },
                        { statusPanel: 'agAggregationComponent' },
                    ],
                },
            };

            window.aa = this.agGrid = new agGrid.Grid(document.querySelector('#myGrid'), gridOptions);
            gridOptions.api.setServerSideDatasource({
                getRows: params => {
                    this.$getJSON("/admin_user/list", JSON.parse(JSON.stringify(params.request)), x => {
                        const lastRow = x.data.total <= params.request.endRow ? x.data.total : -1;
                        params.successCallback(x.data.rows, lastRow);
                    }, x => { params.failCallback() });
                },
            });
            this.agGrid.gridOptions.onRowDoubleClicked = x => this.$router.push(`/admin/user/edit/${x.data.ID}`)
        },
        destroyed() {
            this.agGrid.destroy();
        },
        methods: {
            refresh() {
                this.agGrid.gridOptions.api.purgeServerSideCache()
                this.agGrid.gridOptions.api.deselectAll()
                this.agGrid.gridOptions.api.clearRangeSelection()
            },
            del() {
                const sels = this.agGrid.gridOptions.api.getSelectedRows()
                if (sels.length == 0) return
                let ids = []
                for (let i = 0; i < sels.length; i++) ids.push(sels[i].ID)
                if (!confirm("确定要删除吗？")) return
                this.$getJSON("/admin_user/delete", { IDs: ids.join(",") }, x => {
                    this.refresh()
                });
            }
        },
        filters: {},
        watch: {},
        computed: {},
    }
</script>