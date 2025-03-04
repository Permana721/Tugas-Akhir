import adminDashboard from '../pages/admin/adminDashboard.vue';

import adminBlog from '../pages/admin/blogs/Blog.vue';
import addBlog from '../pages/admin/blogs/addBlog.vue';

import adminFood from '../pages/admin/foods/Food.vue';

import adminRecipe from '../pages/admin/recipes/Recipe.vue';

import adminSport from '../pages/admin/sports/Sport.vue';

import adminUser from '../pages/admin/users/User.vue';

const adminRoutes = [
    { path: '/admin-pg', name: 'adminDashboard', component: adminDashboard, meta: { title: 'Dashboard - Admin IP-CORP' } },

    { path: '/admin-pg/blog', name: 'adminBlog', component: adminBlog, meta: { title: 'Blog - Admin IP-CORP' } },
    { path: '/admin-pg/blog/add-blog', name: 'addBlog', component: addBlog, meta: { title: 'Add Blog - Admin IP-CORP' } },

    { path: '/admin-pg/food', name: 'adminFood', component: adminFood, meta: { title: 'Food - Admin IP-CORP' } },

    { path: '/admin-pg/recipe', name: 'adminRecipe', component: adminRecipe, meta: { title: 'Recipe - Admin IP-CORP' } },

    { path: '/admin-pg/sport', name: 'adminSport', component: adminSport, meta: { title: 'Sport - Admin IP-CORP' } },

    { path: '/admin-pg/user', name: 'adminUser', component: adminUser, meta: { title: 'User - Admin IP-CORP' } },
];

export default adminRoutes;