<template>
    <DefaultLayout>
        <section class="content-header">
            <div class="row">
                <div class="col-xs-12">
                    <div class="box">
                        <div class="box-header">
                            <h3 class="box-title">Data Table With Full Features</h3>
                        </div>
                        <!-- /.box-header -->
                        <div class="box-body">
                            <div id="myGrid" class="ag-theme-balham" style="width: 100%;height:600px;"></div>
                        </div>
                        <!-- /.box-body -->
                    </div>
                    <!-- /.box -->
                </div>
                <!-- /.col -->
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
                    sortable: true,
                    filter: true,
                    filter: 'agTextColumnFilter',
                    filterParams: {
                        filterOptions: ['equals', 'contains']
                    },
                    resizable: true,
                },
                floatingFilter: true,
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
                        { statusPanel: 'agTotalAndFilteredRowCountComponent', align: 'left' },
                        { statusPanel: 'agTotalRowCountComponent', align: 'center' },
                        { statusPanel: 'agFilteredRowCountComponent' },
                        { statusPanel: 'agSelectedRowCountComponent' },
                        { statusPanel: 'agAggregationComponent' },
                    ],
                },
            };

            window.aa = this.agGrid = new agGrid.Grid(document.querySelector('#myGrid'), gridOptions);
            this.agGrid.gridOptions.onCellClicked = x => this.headerName = x.colDef.headerName
            this.refresh()
        },
        destroyed() {
            this.agGrid.destroy();
        },
        methods: {
            refresh() {
                this.$getJSON("/admin_user/list", { startRow: 0, endRow: 100000 }, x => {
                    if (x.data.rows.length == 0) {
                        this.agGrid.gridOptions.api.setRowData([])
                        return toastr.warning("无数据")
                    }
                    this.agGrid.gridOptions.api.setRowData(x.data.rows)
                });
            },
            del() {
                const sels = this.agGrid.gridOptions.api.getSelectedRows()
                if (sels.length == 0) return
                let ids = []
                for (let i = 0; i < sels.length; i++) ids.push(sels[i].ID)
                if (!confirm("确定要删除吗？")) return
                this.$getJSON("/admin_user/delete", { IDs: ids.join(",") }, this.refresh);
            },
        },
        filters: {},
        watch: {},
        computed: {},
    }
</script>