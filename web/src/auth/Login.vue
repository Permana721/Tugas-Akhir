<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-8 bg-white rounded-lg shadow h-screen md:h-auto">
      <div class="flex justify-center items-center">
          <RouterLink to="/">
              <img src="/img/icon/logos.png" alt="Logo" class="w-56 mb-8">
          </RouterLink>
      </div>
      <form>
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700" for="email">Email</label>
          <input
            id="email"
            type="email"
            class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Email"
          />
        </div>
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700" for="password">Password</label>
          <input
            id="password"
            type="password"
            class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="Password"
          />
        </div>
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <input
              id="remember"
              type="checkbox"
              class="w-4 h-4 text-blue-500 border-gray-300 rounded focus:ring-blue-500"
            />
            <label for="remember" class="ml-2 text-sm text-gray-700">Ingat Saya</label>
          </div>
          <a href="#" class="text-sm text-blue-600 hover:underline">Lupa Password?</a>
        </div>
        <button type="submit" class="w-full px-4 py-2 text-white bg-blue-600 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          Masuk
        </button>
      </form>
      <div class="mt-6 text-center">
        <p class="text-sm text-gray-600 mb-4">Atau masuk dengan</p>
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
        Pengguna baru? <RouterLink to="Register" class="text-blue-600 hover:underline">Daftar</RouterLink>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
};
</script>
