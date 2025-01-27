import { createRouter, createWebHistory } from 'vue-router';
import userRoutes from './userRoute';
import authRoutes from './authRoute';

const routes = [
    ...userRoutes,
    ...authRoutes,
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
