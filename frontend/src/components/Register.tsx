import React from "react";
import { Button, Container, Typography, TextField } from "@mui/material";
import { useNavigate } from "react-router-dom";

const Register: React.FC = () => {
    const navigate = useNavigate();

    const handleRegister = () => {
        // 新規登録
        navigate("/quiz");
    };

    return (
        <Container>
            <Typography variant="h4" gutterBottom>
                Register
            </Typography>
            <TextField label="Username" fullWidth margin="normal"></TextField>
            <TextField label="Password" type="password" fullWidth margin="normal"></TextField>
            <TextField label="ConfirmPassword" color="primary" fullWidth margin="normal"></TextField>
            <Button variant="contained" color="primary" onClick={handleRegister}>
                Register
            </Button>
        </Container>
    );
};

export default Register;