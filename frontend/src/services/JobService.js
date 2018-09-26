import axios from 'axios';

class JobService {
  http;
  constructor() {
    this.http = axios.create({
      baseURL: process.env.VUE_APP_JOB_API
    });
  }

  create(job) {
    return this.http.post('/upload', job);
  }
}

export default new JobService();
