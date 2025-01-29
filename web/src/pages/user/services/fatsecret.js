import axios from "axios";
import OAuth from "oauth-1.0a";
import CryptoJS from "crypto-js";

// Konfigurasi OAuth 1.0
const consumerKey = "a65d7594d2f34883978bed61dfc5fc1a";
const consumerSecret = "18c5cb48b9574950964691be5cb54795";

const oauth = OAuth({
  consumer: { key: consumerKey, secret: consumerSecret },
  signature_method: "HMAC-SHA1",
  hash_function(baseString, key) {
    return CryptoJS.HmacSHA1(baseString, key).toString(CryptoJS.enc.Base64);
  },
});

export const searchFood = async (query) => {
  const requestData = {
    url: "https://platform.fatsecret.com/rest/server.api",
    method: "GET",
    data: {
      method: "foods.search",
      format: "json",
      search_expression: query,
    },
  };

  const headers = oauth.toHeader(oauth.authorize(requestData));

  try {
    const response = await axios.get(requestData.url, {
      params: requestData.data,
      headers: headers,
    });
    return response.data.foods.food; // Mengembalikan hasil pencarian
  } catch (error) {
    console.error("Error fetching food data:", error);
    throw error;
  }
};
