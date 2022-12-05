import Vue from 'vue'
import VueRouter from 'vue-router'
const Login = () => import('../views/Login.vue')
const Admin = () => import('../views/Admin.vue')

// 页面路由组件
const Index = () => import('../components/admin/Index.vue')
const Swindler = () => import('../components/swindler/Swindler.vue')
const User = () => import('../components/user/User.vue')

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location,onResolve,onReject) {
  if (onResolve || onReject) return originalPush.call(this,location,onResolve,onReject)
  return originalPush.call(this,location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      title: '请登录'
    }
  },
  {
    path: '/',
    name: 'admin',
    meta: {
      title: 'Query 信息查询页面'
    },
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: 'Query 信息查询页面'
        }
      },
      {
        path: 'swindler',
        component: Swindler,
        meta: {
          title: '骗子信息列表'
        }
      },
      {
        path: 'user',
        component: User,
        meta: {
          title: '用户列表'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/admin') {
    next('login')
  } else {
    next()
  }
})

export default router
