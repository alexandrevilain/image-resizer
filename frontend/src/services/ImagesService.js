import axios from 'axios';

class ImagesService {
  http;
  constructor() {
    this.http = axios.create({
      baseURL: process.env.VUE_APP_IMAGES_API
    });
  }

  getAll() {
    return this.http.get('/images');
  }
}

export default new ImagesService();
