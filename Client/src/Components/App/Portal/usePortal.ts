import axios from "axios";

const UsePortal = () => {
    const fetchRooms = async () => {
        const rooms = await axios.get('/rooms')
            .then(res => {
                return res.data;
            })
            .catch(err => {
                return [];
            })
        if(rooms === null) {
            await createRoom();
            return fetchRooms();
        }
        return rooms;
    } 

    const createRoom = async () => {
        const successful = await axios.post('/rooms')
            .then(res => {
                return res.status === 200;
            })
            .catch(err => {
                return false
            });
        return successful
    }

    const joinRoom = async (roomId: string) => {
        const successful = await axios.post(`/rooms/join/${roomId}`)
            .then(res => {
                return res.status === 200;
            })
            .catch(err => {
                return false
            });

        return successful;
    }
    
    return {
        fetchRooms,
        joinRoom
    }
}

export default UsePortal;