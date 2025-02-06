import adminDashboard from '../pages/admin/adminDashboard.vue';
import Blog from '../pages/admin/blogs/Blog.vue';
import Food from '../pages/admin/foods/Food.vue';
import Recipe from '../pages/admin/recipes/Recipe.vue';
import Sport from '../pages/admin/sports/Sport.vue';
import User from '../pages/admin/users/User.vue';

const adminRoutes = [
    { path: '/admin-pg', name: 'adminDashboard', component: adminDashboard, meta: { title: 'Dashboard - Admin IP-CORP' } },

    { path: '/admin-pg/blog', name: 'Blog', component: Blog, meta: { title: 'Blog - Admin IP-CORP' } },

    { path: '/admin-pg/food', name: 'Food', component: Food, meta: { title: 'Food - Admin IP-CORP' } },

    { path: '/admin-pg/recipe', name: 'Recipe', component: Recipe, meta: { title: 'Recipe - Admin IP-CORP' } },

    { path: '/admin-pg/sport', name: 'Sport', component: Sport, meta: { title: 'Sport - Admin IP-CORP' } },

    { path: '/admin-pg/user', name: 'User', component: User, meta: { title: 'User - Admin IP-CORP' } },
];

export default adminRoutes;