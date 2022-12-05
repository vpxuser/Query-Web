import Vue from 'vue'
import VueRouter from 'vue-router'
const Guest = () => import('../views/Guest.vue')

// 页面路由组件
const Index = () => import('../components/guest/Index.vue')
const Swindler = () => import('../components/swindler/Swindler.vue')

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location,onResolve,onReject) {
  if (onResolve || onReject) return originalPush.call(this,location,onResolve,onReject)
  return originalPush.call(this,location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'guest',
    meta: {
      title: 'Query 信息查询页面'
    },
    component: Guest,
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
})

export default router
