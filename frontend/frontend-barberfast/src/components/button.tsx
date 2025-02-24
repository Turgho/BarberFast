import React from "react";

interface ButtonProps {
    text: string;
    onClick?: () => void;
    variant?: "primary" | "secondary";
    size?: "sm" | "md" | "lg";
    disabled?: boolean;
}

export default function Button({ text, onClick, variant = "primary", size = "lg", disabled = false }: ButtonProps) {
    const baseStyle = "rounded-full font-julius text-center transition-all duration-300 focus:outline-none transform";

    // Tamanhos de botão com responsividade
    const sizeStyles = {
        sm: "px-4 py-2 text-sm",
        md: "px-6 py-3 text-base",
        lg: "px-8 py-4 text-lg md:px-12 md:py-6 md:text-xl lg:px-16 lg:py-8 lg:text-2xl", // Ajustes maiores em telas grandes
    };

    // Estilos de variante para o botão
    const variantStyles = {
        primary: "bg-laranja text-white hover:bg-gray-600 focus:ring-2 focus:ring-offset-2 focus:ring-laranja disabled:bg-gray-300 disabled:cursor-not-allowed hover:scale-105 active:scale-95",
        secondary: "bg-black text-white hover:bg-gray-100 focus:ring-2 focus:ring-offset-2 focus:ring-gray-300 disabled:bg-gray-300 disabled:cursor-not-allowed hover:scale-105 active:scale-95",
    };

    return (
        <button 
            onClick={onClick} 
            className={`${baseStyle} ${sizeStyles[size]} ${variantStyles[variant]}`}
            disabled={disabled}
        >
            {text}
        </button>
    );
}
