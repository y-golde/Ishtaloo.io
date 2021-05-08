const UseLoginModal = () => {
    const handleLoginClick = (userName: string) => {
        alert(userName)
    }

    return {
        handleLoginClick
    };
}

export default UseLoginModal;