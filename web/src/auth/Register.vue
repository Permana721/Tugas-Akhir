<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
        <div class="w-full max-w-md p-8 bg-white rounded-lg shadow">
            <h1 class="text-2xl font-bold text-center mb-6">FUFUFAFA</h1>
            <p class="text-center text-gray-600 mb-6">Create your account</p>
            <form>
                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700" for="name">Full Name</label>
                    <input
                    id="name"
                    type="text"
                    class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
                    placeholder="Full Name"
                    />
                </div>
                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700" for="email">Email</label>
                    <input
                    id="email"
                    type="email"
                    class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
                    placeholder="Email"
                    />
                </div>
                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700" for="password">Password</label>
                    <input
                    id="password"
                    type="password"
                    class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
                    placeholder="Password"
                    />
                </div>
                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700" for="confirm-password">Confirm Password</label>
                    <input
                    id="confirm-password"
                    type="password"
                    class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
                    placeholder="Confirm Password"
                    />
                </div>
                <button
                    type="submit"
                    class="w-full px-4 py-2 text-white bg-purple-600 rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500"
                >
                    Register
                </button>
            </form>
            <div class="mt-6 text-center">
                <p class="text-sm text-gray-600 mb-4">Or sign in with</p>
                <div class="flex justify-center space-x-4">
                <button class="p-2 bg-gray-100 rounded-lg shadow hover:bg-gray-200 focus:outline-none">
                    <img src="/img/icon/facebook.png" alt="Facebook" class="w-6 h-6" />
                </button>
                <button class="p-2 bg-gray-100 rounded-lg shadow hover:bg-gray-200 focus:outline-none">
                    <img src="/img/icon/instagram.png" alt="Instagram" class="w-6 h-6" />
                </button>
                <button class="p-2 bg-gray-100 rounded-lg shadow hover:bg-gray-200 focus:outline-none">
                    <img src="/img/icon/google.png" alt="Google" class="w-6 h-6" />
                </button>
                </div>
            </div>
            <p class="mt-6 text-sm text-center text-gray-600">
                Already have an account? <RouterLink to="/login" class="text-purple-600 hover:underline">Sign in</RouterLink>
            </p>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Register",
    };
</script>