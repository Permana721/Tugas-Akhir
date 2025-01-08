import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../pages/DashboardPage.vue';
import Login from '../auth/Login.vue';
import Register from '../auth/Register.vue';

const routes = [
    { path: '/', redirect: '/dashboard' }, 
    { path: '/dashboard', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard' } },
    { path: '/login', name: 'Login', component: Login, meta: { title: 'Login' } },
    { path: '/register', name: 'Register', component: Register, meta: { title: 'Register' } },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title;
    } else {
        document.title = 'Default Title';
    }

    next();
});

export default router;