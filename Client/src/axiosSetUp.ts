import axios from 'axios';


const setUp = () => {
    axios.defaults.baseURL = globalThis.API_URL;
}

export default setUp;