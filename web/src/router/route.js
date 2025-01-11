import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../pages/DashboardPage.vue';
import Login from '../auth/Login.vue';
import Blog from '../pages/Blog/BlogPage.vue';
import Diet from '../pages/Diet/DietPage.vue';
import Plan from '../pages/Plan/PlanPage.vue';
import Food from '../pages/Food/FoodPage.vue';
import Recipe from '../pages/Recipe/RecipePage.vue';
import Register from '../auth/Register.vue';

const routes = [
    // { path: '/', redirect: '/dashboard' }, 
    { path: '/', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard - FUFUFAFA' } },
    { path: '/blog', name: 'Blog', component: Blog, meta: { title: 'Blog' } },
    { path: '/diet', name: 'Diet', component: Diet, meta: { title: 'Diet' } },
    { path: '/plan', name: 'Plan', component: Plan, meta: { title: 'Rencana' } },
    { path: '/food', name: 'Food', component: Food, meta: { title: 'Makanan' } },
    { path: '/recipe', name: 'Recipe', component: Recipe, meta: { title: 'Recipe' } },

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