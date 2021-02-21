import axios from 'axios';

const api = axios.create({
  baseURL: 'https://vh7.uk/',
  timeout: 1000,
});

async function getLanguages() {
  const response = await api.get('/languages');
  return response.data;
}

async function createShorten(url: string) {
  const response = await api.post('/shorten', {
    url,
  });
  return response.data;
}

async function createPaste(language: string, code: string) {
  const response = await api.post('/pastet', {
    language,
    code,
  });
  return response.data;
}

// async function createUpload(file: string) {

// }

export default {
  getLanguages,
  createShorten,
  createPaste,
  // createUpload
};
