import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

const router = new Router(
    {
        mode: 'history',
        routes: [
            {
                path: '/',  // 程序启动默认路由
                component: () => import('../components/Whole.vue'),
                meta: {title: '首页'},
                children: [
                    {
                        path: '/uploadPic',
                        component: () => import('../components/UploadPic.vue'),
                        meta: {title: '上传图片'}
                    },
                    {
                        path: '/uploadVideo',
                        component: () => import('../components/UploadVideo.vue'),
                        meta: {title: '上传视频'}
                    },
                ]
            },
        ]
    }
);


export default router
