import axios, {AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse} from 'axios';

// 创建实例
const axiosInstance: AxiosInstance = axios.create({
    // 前缀
    baseURL: import.meta.env.VITE_API_BASEURL,
    // 超时
    timeout: 1000 * 30,
    // 请求头
    headers: {
        'Content-Type': 'application/json',
    },
});

// 请求拦截器
axiosInstance.interceptors.request.use(
    (config: AxiosRequestConfig) => {
        return config;
    },
    (error: AxiosError) => {
        return Promise.reject(error);
    },
);

// 响应拦截器
axiosInstance.interceptors.response.use(
    (response: AxiosResponse) => {
        const {data} = response;
        if (data.code === 200) {
            return data.data;
        }
        return Promise.reject(data.msg);
    },
    (error: AxiosError) => {
        return Promise.reject(error);
    },
);
const service = {
    get<T = unknown>(url: string, data?: object): Promise<T> {
        return axiosInstance.get(url, {params: data});
    },

    post<T = unknown>(url: string, data?: object): Promise<T> {
        return axiosInstance.post(url, data);
    },

    put<T = unknown>(url: string, data?: object): Promise<T> {
        return axiosInstance.put(url, data);
    },

    delete<T = unknown>(url: string, data?: object): Promise<T> {
        return axiosInstance.delete(url, data);
    },
    upload: (url: string, file: FormData | File) =>
        axiosInstance.post(url, file, {
            headers: {'Content-Type': 'multipart/form-data'},
        }),
};

export default service;
