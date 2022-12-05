<template>
  <div>
    <a-card>
      <a-row :gutter='20'>
        <a-col span='6'>
          <a-input-search v-model='queryParam.tieba' placeholder='请输入贴吧ID查找' enter-button allowClear @search='searchSwindler' />
        </a-col>
      </a-row>

      <a-table rowKey='ID' :columns='columns' :pagination='pagination' :dataSource='swindlerlist' bordered @change='handleTableChange'>
      </a-table>
    </a-card>
  </div>
</template>

<script>
import day from 'dayjs'

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '创建时间',
    dataIndex: 'CreatedAt',
    width: '20%',
    key: 'CreatedAt',
    align: 'center',
    customRender: (val) => {
      return val ? day(val).format('YYYY年MM月DD日 HH:mm') : '暂无'
    }
  },
  {
    title: '贴吧号',
    dataIndex: 'tieba',
    width: '20%',
    key: 'tieba',
    align: 'center'
  },
  {
    title: '微信号',
    dataIndex: 'wechat',
    width: '20%',
    key: 'wechat',
    align: 'center'
  },
  {
    title: '手机号',
    dataIndex: 'phone',
    width: '20%',
    key: 'phone',
    align: 'center'
  }
]

export default {
  data() {
    return {
      pagination:{
        pageSizeOptions: ['5','10','20'],
        pageSize: 5,
        current: 1,
        total: 0,
        showSizeChanger: true,
        showQuickJumper: true,
        hideOnSinglePage: false,
        showTotal: (total) => `共${total}条`,
        onChange: (page, pageSize) => {
          this.pagination.current = page
          this.pagination.pageSize = pageSize
        }
      },
      swindlerlist: [],
      queryParam: {
        tieba: '',
        pagesize: 5,
        pagenum: 1
      },
      columns,
      editVisible: false,
      addSwindlerVisible: false,
      editSwindlerVisible: false,
    }
  },
  created () {
    this.searchSwindler()
  },
  methods: {
    // 按名称查询用户
    async searchSwindler() {
      const { data: res } = await this.$http.get('swindler/list',{
        params: {
          tieba: this.queryParam.tieba,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.swindlerlist = res.data
      this.pagination.total = res.total
    },

    handleTableChange(pagination) {
      this.pagination = pagination
      this.queryParam.pagenum = pagination.current
      this.queryParam.pagesize = pagination.pageSize
      this.searchSwindler()
    },
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
