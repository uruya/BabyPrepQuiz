import React from "react";
import { Button, Container, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";

const Welcome: React.FC = () => {
    const navigate = useNavigate();

    const handleLogin = () => {
        navigate("/login");
    };

    const handleRegister = () => {
        navigate("/register");
    }

    const handleContinueWithoutLogin = () => {
        navigate("/quiz");
    };

    return (
        <Container>
            <Typography variant="h4" gutterBottom>
                Welcome to the Quiz App
            </Typography>
            <Button variant="contained" color="primary" onClick={handleLogin}>
                Log In
            </Button>
            <Button variant="contained" color="secondary" onClick={handleRegister}>
                Register
            </Button>
            <Button variant="outlined" color="info" onClick={handleContinueWithoutLogin}>
                Continue with out Login
            </Button>
        </Container>
    );
};

export default Welcome;