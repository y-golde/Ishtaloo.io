const UseLoginModal = () => {
    const handleLoginClick = (userName: string) => {
        alert(userName)
    }

    const shouldDisable = (userName: string) => {
        return userName === ''
    }

    return {
        handleLoginClick,
        shouldDisable
    };
}

export default UseLoginModal;