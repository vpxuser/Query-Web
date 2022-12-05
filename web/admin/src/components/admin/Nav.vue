<template>
  <a-layout-sider breakpoint='lg' v-model='collapsed'>
    <div class='log'>
      <span>{{ collapsed ? 'Query' : 'My Query' }}</span>
    </div>

    <a-menu theme='dark' mode='inline' @click='goToPage'>
      <a-menu-item key='swindler'>
        <a-icon type='book' />
          <span>骗子信息列表</span>
      </a-menu-item>
      <a-menu-item key='user'>
        <a-icon type='user' />
        <span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  methods: {
    goToPage (item) {
      this.$router.push(item.key).catch((err) => err)
    }
  }
}
</script>

<style scoped>
.log{
  height: 32px;
  margin: 16px;
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>
