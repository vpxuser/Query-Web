<template>
  <div>
    <a-card>
      <a-row :gutter='20'>
        <a-col span='6'>
          <a-input-search v-model='queryParam.username' placeholder='请输入用户名查找' enter-button allowClear @search='searchUser' />
        </a-col>
        <a-col span='4'>
          <a-button type='primary' icon='plus-circle' style='margin-right: 15px' @click='addUserVisible = true'>新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey='ID' :columns='columns' :pagination='pagination' :dataSource='userlist' bordered @change='handleTableChange'>
        <template slot='action' slot-scope='data'>
          <div class='actionSlot'>
            <a-button type='primary' icon='edit' style='margin-right: 15px' @click='editUser(data.ID)'>编辑</a-button>
            <a-button type='danger' icon='delete' style='margin-right: 15px' @click='deleteUser(data.ID)'>删除</a-button>
            <a-button type='info' icon='info-circle' style='margin-right: 15px' @click='rePassword(data.ID)'>修改密码</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增用户 -->
    <a-modal closable destroyOnClose title='新增用户' :visible='addUserVisible' width='60%' @ok='addUserOk' @cancel='addUserCancel'>
      <a-form-model :model='addUserModel' :rules='addUserRules' ref='addUserRef'>
        <a-form-model-item label='用户名' prop='username'>
          <a-input v-model='addUserModel.username'></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label='密码' prop='password'>
          <a-input-password v-model='addUserModel.password'></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label='确认密码' prop='checkpass'>
          <a-input-password v-model='addUserModel.checkpass'></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑用户 -->
    <a-modal closable destroyOnClose title='编辑用户' :visible='editUserVisible' width='60%' @ok='editUserOk' @cancel='editUserCancel'>
      <a-form-model :model='editUserModel' :rules='editUserRules' ref='editUserRef'>
        <a-form-model-item label='用户名' prop='username'>
          <a-input v-model='editUserModel.username'></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 修改密码 -->
    <a-modal closable destroyOnClose title='修改密码' :visible='rePasswordVisible' width='60%' @ok='rePasswordOk' @cancel='rePasswordCancel'>
      <a-form-model :model='rePasswordModel' :rules='rePasswordRules' ref='rePasswordRef'>
        <a-form-model-item has-feedback label='密码' prop='password'>
          <a-input-password v-model='rePasswordModel.password'></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label='确认密码' prop='checkpass'>
          <a-input-password v-model='rePasswordModel.checkpass'></a-input-password>
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
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '注册时间',
    dataIndex: 'CreatedAt',
    width: '20%',
    key: 'CreatedAt',
    align: 'center',
    customRender: (val) => {
      return val ? day(val).format('YYYY年MM月DD日 HH:mm') : '暂无'
    }
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
      userlist: [],
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      addUserModel: {
        username: '',
        password:'',
        checkPass: ''
      },
      editUserModel: {
        id: 0,
        username: ''
      },
      rePasswordModel: {
        id: 0,
        password: '',
        checkpass: ''
      },
      columns,
      editVisible: false,
      addUserRules: {
        username: [
          {
            validator: (rule,value,callback) => {
              if (this.addUserModel.username === '') {
                callback(new Error('请输入用户名'))
              }
              if ([...this.addUserModel.username].length < 4 || [...this.addUserModel.username].length > 12) {
                callback(new Error('用户名应该大于4个字符并且小于12个字符'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule,value,callback) => {
              if (this.addUserModel.password === '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.addUserModel.password].length < 8 || [...this.addUserModel.password].length > 20) {
                callback(new Error('密码应该大于8个字符并且小于20个字符'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule,value,callback) => {
              if (this.addUserModel.checkpass === '') {
                callback(new Error('请输入密码'))
              }
              if (this.addUserModel.password !== this.addUserModel.checkpass) {
                callback(new Error('密码不一致请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editUserRules: {
        username: [
          {
            validator: (rule,value,callback) => {
              if (this.editUserModel.username === '') {
                callback(new Error('请输入用户名'))
              }
              if ([...this.editUserModel.username].length < 4 || [...this.editUserModel.username].length > 12) {
                callback(new Error('用户名应该大于4个字符并且小于12个字符'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      rePasswordRules: {
        password: [
          {
            validator: (rule,value,callback) => {
              if (this.rePasswordModel.password === '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.rePasswordModel.password].length < 8 || [...this.rePasswordModel.password].length > 20) {
                callback(new Error('密码应该大于8个字符并且小于20个字符'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule,value,callback) => {
              if (this.rePasswordModel.checkPass === '') {
                callback(new Error('请输入密码'))
              }
              if (this.rePasswordModel.password !== this.rePasswordModel.checkpass) {
                callback(new Error('密码不一致请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addUserVisible: false,
      editUserVisible: false,
      rePasswordVisible: false
    }
  },
  created () {
    this.searchUser()
  },
  methods: {
    // 按名称查询用户
    async searchUser() {
      const { data: res } = await this.$http.get('user/list',{
        params: {
          username: this.queryParam.username,
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
      this.userlist = res.data
      this.pagination.total = res.total
    },

    handleTableChange(pagination) {
      this.pagination = pagination
      this.queryParam.pagenum = pagination.current
      this.queryParam.pagesize = pagination.pageSize
      this.searchUser()
    },

    // 添加用户
    addUserOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('user/add',{
          username: this.addUserModel.username,
          password: this.addUserModel.password,
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addUserRef.resetFields()
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        await this.searchUser()
      })
    },

    // 取消添加用户操作
    addUserCancel() {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
      this.$message.info('新增用户已取消')
    },

    // 删除用户
    deleteUser(id) {
      this.$confirm({
        title: '提示',
        content: '确认删除此用户吗？一旦删除，无法恢复！',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          await this.searchUser()
        },
        onCancel: () => {
          this.$message.info('已取消删除操作')
        }
      })
    },

    // 编辑用户
    async editUser(id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.editUserModel = res.data
      this.editUserModel.id = id
    },
    editUserOk() {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`user/${this.editUserModel.id}`,{
          username: this.editUserModel.username
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.editUserRef.resetFields()
        this.editUserVisible = false
        this.$message.success('更新用户信息成功')
        await this.searchUser()
      })
    },

    // 取消编辑用户操作
    editUserCancel() {
      this.$refs.editUserRef.resetFields()
      this.editUserVisible = false
      this.$message.info('编辑已取消')
    },

    //修改密码
    async rePassword(id) {
      this.rePasswordVisible = true
      //const { data: res } = await this.$http.get(`user/${id}`)
      //this.rePasswordModel = res.data
      this.rePasswordModel.id = id
    },
    rePasswordOk() {
      this.$refs.rePasswordRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`repassword/${this.rePasswordModel.id}`,{
          password: this.rePasswordModel.password
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.rePasswordVisible = false
        this.$message.success('修改密码成功')
        await this.searchUser()
      })
    },
    rePasswordCancel() {
      this.$refs.rePasswordRef.resetFields()
      this.rePasswordVisible = false
      this.$message.info('已取消')
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
