import Home from "../views/Home.vue";
import AboutView from "../views/AboutView.vue";
import ChatView from "../views/ChatView.vue";
import MainPage from "../views/MainPage.vue";

const routes = [
    { path: '/', component: Home ,
        children: [
            {
                name: 'home',
                path: 'home/',
                component: MainPage
            },
            {
                name: 'about',
                path: 'about/',
                component: AboutView
            },
            {
                name:'chat',
                path:'chat/',
                component: ChatView,
            },
            {
                name :'countDown',
                path :'countDown/',
                component : () => import('../views/CountdownView.vue')
            },

        ]
    },
    { path: '/about', component: AboutView},
    { path: '/login', component: () => import('../views/LoginView.vue') ,name : 'login'},
]


export default routes;