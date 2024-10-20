import axios, { AxiosError, AxiosInstance, AxiosResponse } from 'axios';

// 创建实例
const axiosInstance: AxiosInstance = axios.create({
  // 前缀
  baseURL: import.meta.env.VITE_MEMPOOL,
  // 超时
  timeout: 1000 * 30,
  // 请求头
  headers: {
    'Content-Type': 'application/json',
  },
});

axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data } = response;
    return data;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  },
);

export const getFeeRate = (): Promise<any> => {
  return axiosInstance.get(`/v1/fees/recommended`);
};

export const getTipHeight = (): Promise<any> => {
  return axiosInstance.get(`/blocks/tip/height`);
};
