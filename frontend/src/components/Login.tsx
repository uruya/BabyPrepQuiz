import React from "react";
import { Button, Container, Typography, TextField } from "@mui/material";
import { useNavigate } from "react-router-dom";

const Login: React.FC = () => {
    const navigate = useNavigate();

    const handleLogin = () => {
        // ログイン処理
        navigate("/quiz");
    };

    return (
        <Container>
            <Typography variant="h4" gutterBottom>
                Log In
            </Typography>
            <TextField label="Username" fullWidth margin="normal"></TextField>
            <TextField label="Password" type="password" fullWidth margin="normal"></TextField>
            <Button variant="contained" color="primary" onClick={handleLogin}>
                Log In
            </Button>
        </Container>
    );
};

export default Login;