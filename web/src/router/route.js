import { createRouter, createWebHistory } from 'vue-router';
import userRoutes from './userRoute';
import authRoutes from './authRoute';
import adminRoutes from './adminRoute';

const routes = [
    ...userRoutes,
    ...authRoutes,
    ...adminRoutes,
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

console.log('Registered routes:', routes);

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title;
    } else {
        document.title = 'Default Title';
    }
    next();
});

export default router;