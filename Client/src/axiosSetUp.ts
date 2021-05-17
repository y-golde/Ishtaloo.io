import axios from 'axios';

const setUp = () => {
	axios.defaults.withCredentials = true;
	axios.defaults.baseURL = globalThis.API_URL;
};

export default setUp;
