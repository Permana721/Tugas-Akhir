import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../pages/DashboardPage.vue';
// import Login from '../pages/LoginPage.vue';

const routes = [
    { path: '/', redirect: '/dashboard' }, // Redirect ke dashboard langsung
    { path: '/dashboard', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard' } },
    // { path: '/login', name: 'Login', component: Login, meta: { title: 'Login' } },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    // Set title untuk setiap halaman
    if (to.meta.title) {
        document.title = to.meta.title;
    } else {
        document.title = 'Default Title';
    }

    next(); // Hilangkan cek login agar semua halaman bisa diakses
});

export default router;
