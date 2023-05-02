import axios from "axios";

export default axios.create({
    baseURL: `https://api.imgbb.com/1/upload?key=${import.meta.env.VITE_IMAGE_API}`
})