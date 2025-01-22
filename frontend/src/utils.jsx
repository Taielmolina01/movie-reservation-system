export const CinemaChainName = 'Starlight Cinemas';
const BACK_URL = import.meta.env.BACKEND_URL;

const HandleSignIn = async(email, password, setLoading, setError, navigate) => {
    try {
        setLoading(true);
        const res = await fetch(`${BACK_URL}/login`, {
            method: 'POST',
            body: new URLSearchParams({
                username: email,
                password: password
            })
        });

        const data = await res.json();

        if (!res.ok) {
            setError(data.detail)
        } else {
            localStorage.setItem('accessToken', data.access_token);
            localStorage.setItem('user'), JSON.stringify(data.user);
            {/* navigate("/home")*/}
        }
    } catch(error) {
        setError(error)
        return;
    } finally {
        setLoading(false);
    }
}