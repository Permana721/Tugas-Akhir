import adminDashboard from '../pages/admin/adminDashboard.vue';

const adminRoutes = [
    { path: '/admin-pg', name: 'adminDashboard', component: adminDashboard, meta: { title: 'Admin Dashboard - IP-CORP' } },
];

export default adminRoutes;