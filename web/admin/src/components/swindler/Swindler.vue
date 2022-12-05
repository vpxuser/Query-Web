<template>
  <div>
    <a-card>
      <a-row :gutter='20'>
        <a-col span='6'>
          <a-input-search v-model='queryParam.tieba' placeholder='请输入贴吧ID查找' enter-button allowClear @search='searchSwindler' />
        </a-col>
        <a-col span='4'>
          <a-button type='primary' icon='plus-circle' style='margin-right: 15px' @click='addSwindlerVisible = true'>新增</a-button>
        </a-col>
      </a-row>

      <a-table rowKey='ID' :columns='columns' :pagination='pagination' :dataSource='swindlerlist' bordered @change='handleTableChange'>
        <template slot='action' slot-scope='data'>
          <div class='actionSlot'>
            <a-button type='primary' icon='edit' style='margin-right: 15px' @click='editSwindler(data.ID)'>编辑</a-button>
            <a-button type='danger' icon='delete' style='margin-right: 15px' @click='deleteSwindler(data.ID)'>删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增用户 -->
    <a-modal closable destroyOnClose title='新增骗子信息' :visible='addSwindlerVisible' width='60%' @ok='addSwindlerOk' @cancel='addSwindlerCancel'>
      <a-form-model :model='addSwindlerModel' :rules='addSwindlerRules' ref='addSwindlerRef'>
        <a-form-model-item label='贴吧号' prop='tieba'>
          <a-input v-model='addSwindlerModel.tieba'></a-input>
        </a-form-model-item>
        <a-form-model-item label='微信号' prop='wechat'>
          <a-input v-model='addSwindlerModel.wechat'></a-input>
        </a-form-model-item>
        <a-form-model-item label='手机号' prop='phone'>
          <a-input v-model='addSwindlerModel.phone'></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑骗子信息 -->
    <a-modal closable destroyOnClose title='编辑骗子信息' :visible='editSwindlerVisible' width='60%' @ok='editSwindlerOk' @cancel='editSwindlerCancel'>
      <a-form-model :model='editSwindlerModel' :rules='editSwindlerRules' ref='editSwindlerRef'>
        <a-form-model-item label='微信号' prop='wechat'>
          <a-input v-model='editSwindlerModel.wechat'></a-input>
        </a-form-model-item>
        <a-form-model-item label='手机号' prop='phone'>
          <a-input v-model='editSwindlerModel.phone'></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
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
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    scopedSlots: {customRender:'action'},
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
      addSwindlerModel: {
        tieba: '',
        wechat: '',
        phone: '',
      },
      editSwindlerModel: {
        id: 0,
        wechat: '',
        phone: '',
      },
      columns,
      editVisible: false,
      addSwindlerRules: {
        tieba: [
          {
            validator: (rule,value,callback) => {
              if (this.addSwindlerModel.tieba === '') {
                callback(new Error('请输入贴吧ID'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        wechat: [
          {
            validator: (rule,value,callback) => {
              if (this.addSwindlerModel.wechat !== '' && (this.addSwindlerModel.wechat.length < 6 || this.addSwindlerModel.wechat.length > 28)) {
                callback(new Error('微信号长度不能小于6个字符且不能大于28个字符'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        phone: [
          {
            validator: (rule,value,callback) => {
              if (this.addSwindlerModel.phone !== '' && this.addSwindlerModel.phone.length !== 11) {
                callback(new Error('电话号码长度应该为11个数字'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editSwindlerRules: {
        wechat: [
          {
            validator: (rule,value,callback) => {
              if (this.addSwindlerModel.wechat !== '' && (this.addSwindlerModel.wechat.length < 6 || this.addSwindlerModel.wechat.length > 28)) {
                callback(new Error('微信号长度不能小于6个字符且不能大于28个字符'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        phone: [
          {
            validator: (rule,value,callback) => {
              if (this.addSwindlerModel.phone !== '' && this.addSwindlerModel.phone.length !== 11) {
                callback(new Error('电话号码长度应该为11个数字'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
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
      if ( res.status !== 200 ) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          await this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.swindlerlist = res.data
      this.pagination.total = res.total
    },

    handleTableChange(pagination) {
      this.pagination = pagination
      this.queryParam.pagenum = pagination.current
      this.queryParam.pagesize = pagination.pageSize
      this.searchSwindler()
    },

    // 添加用户
    addSwindlerOk() {
      this.$refs.addSwindlerRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('swindler/add',{
          tieba: this.addSwindlerModel.tieba,
          wechat: this.addSwindlerModel.wechat,
          phone: this.addSwindlerModel.phone
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addSwindlerRef.resetFields()
        this.addSwindlerVisible = false
        this.$message.success('添加骗子信息成功')
        await this.searchSwindler()
      })
    },

    // 取消添加用户操作
    addSwindlerCancel() {
      this.$refs.addSwindlerRef.resetFields()
      this.addSwindlerVisible = false
      this.$message.info('新增骗子信息已取消')
    },

    // 删除用户
    deleteSwindler(id) {
      this.$confirm({
        title: '提示',
        content: '确认删除此骗子信息吗？一旦删除，无法恢复！',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`swindler/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          await this.searchSwindler()
        },
        onCancel: () => {
          this.$message.info('已取消删除操作')
        }
      })
    },

    // 编辑用户
    async editSwindler(id) {
      this.editSwindlerVisible = true
      const { data: res } = await this.$http.get(`swindler/${id}`)
      this.editSwindlerModel = res.data
      this.editSwindlerModel.id = id
    },
    editSwindlerOk() {
      this.$refs.editSwindlerRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`swindler/${this.editSwindlerModel.id}`,{
          wechat: this.editSwindlerModel.wechat,
          phone: this.editSwindlerModel.phone
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.editSwindlerRef.resetFields()
        this.editSwindlerVisible = false
        this.$message.success('更新骗子信息成功')
        await this.searchSwindler()
      })
    },

    // 取消编辑用户操作
    editSwindlerCancel() {
      this.$refs.editSwindlerRef.resetFields()
      this.editSwindlerVisible = false
      this.$message.info('编辑已取消')
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
