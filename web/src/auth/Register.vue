<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
      <div class="w-full max-w-md p-8 bg-white rounded-lg shadow">
        <h1 class="text-2xl font-bold text-center mb-6">FUFUFAFA</h1>
  
        <!-- Rincian Saya -->
        <h2 class="text-lg font-semibold text-gray-700 mb-4">Rincian Saya</h2>
        <form>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="gender">Jenis Kelamin</label>
            <select
              id="gender"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option>Laki-laki</option>
              <option>Perempuan</option>
            </select>
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="weight">Berat Badan Saat Ini (kg)</label>
            <input
              id="weight"
              type="number"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Berat Badan"
            />
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="height">Tinggi (cm)</label>
            <input
              id="height"
              type="number"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Tinggi"
            />
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="birthdate">Tanggal Lahir</label>
            <input
              id="birthdate"
              type="date"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="activity-level">Tingkat Aktivitas*</label>
            <select
              id="activity-level"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              <option>Jarang Sekali</option>
              <option>Sedikit Aktif</option>
              <option>Aktif</option>
              <option>Sangat Aktif</option>
            </select>
          </div>
  
          <!-- Tujuan Saya -->
          <h2 class="text-lg font-semibold text-gray-700 mb-4 mt-8">Tujuan Saya</h2>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="diet-goal">Tujuan Diet</label>
            <input
              id="diet-goal"
              type="text"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Contoh: Menurunkan berat badan"
            />
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="weight-goal">Tujuan Berat Badan (kg)</label>
            <input
              id="weight-goal"
              type="number"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Tujuan Berat Badan"
            />
          </div>
  
          <!-- Akun FUFUFAFA -->
          <h2 class="text-lg font-semibold text-gray-700 mb-4 mt-8">Akun FUFUFAFA</h2>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="username">Nama Anggota</label>
            <input
              id="username"
              type="text"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Nama Anggota"
            />
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="email">Alamat Email</label>
            <input
              id="email"
              type="email"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Email"
            />
          </div>
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700" for="password">Kata Sandi</label>
            <input
              id="password"
              type="password"
              class="w-full px-4 py-2 mt-1 border rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Kata Sandi"
            />
          </div>
  
          <!-- Submit Button -->
          <button
            type="submit"
            class="w-full px-4 py-2 text-white bg-purple-600 rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500"
          >
            Register
          </button>
        </form>
  
        <!-- Social Media Buttons -->
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
  
        <!-- Already have an account -->
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
  