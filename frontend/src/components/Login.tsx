import React, { useState } from "react";
import { Button, Container, Typography, TextField } from "@mui/material";
import { useNavigate } from "react-router-dom";

const Login: React.FC = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async () => {
        try {
            // ログイン処理
            const response = await fetch("http://localhost:8080/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username: username, password: password }),
            });
    
            if (response.ok) {
                navigate("/quiz");
            } else {
                const errorData = await response.json();
                setError(errorData.message || "Login failed");
            }
        } catch (error) {
            setError("An error occurred during login");
        }
    };

    return (
        <Container>
            <Typography variant="h4" gutterBottom>
                Log In
            </Typography>
            <TextField label="Username" fullWidth margin="normal" value={username} onChange={(e) => setUsername(e.target.value)}></TextField>
            <TextField label="Password" type="password" fullWidth margin="normal" value={password} onChange={(e) => setPassword(e.target.value)}></TextField>
            {error && <Typography color="error">{error}</Typography>}
            <Button variant="contained" color="primary" onClick={handleLogin}>
                Log In
            </Button>
        </Container>
    );
};

export default Login;