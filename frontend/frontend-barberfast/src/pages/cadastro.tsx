import Navbar from "@/components/navbar";
import "@/styles/globals.css";
import Link from "next/link";
import { useState } from "react";
import { useRouter } from 'next/router'; // Importando o useRouter

// Componente de Cadastro
export default function Cadastro() {
    // Estado para armazenar os dados do formulário
    const [formData, setFormData] = useState({
        nome: "",
        email: "",
        telefone: "",
        senha: "",
        confirmarSenha: "",
        termos: false,
    });

    // Estado para armazenar os erros de validação
    const [errors, setErrors] = useState<{ [key: string]: string }>({});
    const [isSuccess, setIsSuccess] = useState(false); // Novo estado para a mensagem de sucesso
    const [isSubmitting, setIsSubmitting] = useState(false); // Estado para controle do botão
    const router = useRouter(); // Instanciando o router

    // Função para atualizar os dados no estado conforme o usuário digita
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value, type, checked } = e.target;
        setFormData({ 
            ...formData, 
            [name]: type === "checkbox" ? checked : value 
        });
    };

    // Função para formatar o número de telefone automaticamente
    const formatPhoneNumber = (value: string) => {
        // Remove tudo que não for número
        value = value.replace(/\D/g, "");

        // Aplica a máscara (xx) xxxxx-xxxx
        if (value.length <= 2) {
            return `(${value}`;
        }
        if (value.length <= 6) {
            return `(${value.slice(0, 2)}) ${value.slice(2)}`;
        }
        return `(${value.slice(0, 2)}) ${value.slice(2, 7)}-${value.slice(7, 11)}`;
    };

    // Função para lidar com a alteração do telefone e aplicar a máscara automaticamente
    const handlePhoneChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: formatPhoneNumber(value),
        });
    };

    // Função para validar o formulário e lidar com o envio
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        let validationErrors: { [key: string]: string } = {};
    
        // Validações...
        if (!formData.nome) validationErrors.nome = "O nome é obrigatório.";
        if (!formData.email.includes("@")) validationErrors.email = "Email inválido.";
        if (!formData.telefone.match(/^\(\d{2}\) \d{5}-\d{4}$/)) {
            validationErrors.telefone = "O telefone deve ser no formato (xx) xxxxx-xxxx.";
        }
        if (!formData.senha.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{6,}$/)) {
            validationErrors.senha = "A senha deve ter no mínimo 6 caracteres, incluindo maiúsculas, minúsculas, números e caracteres especiais.";
        }
        if (formData.senha !== formData.confirmarSenha) validationErrors.confirmarSenha = "As senhas não coincidem.";
        if (!formData.termos) validationErrors.termos = "Você precisa aceitar os termos.";
    
        setErrors(validationErrors);
    
        // Se houver erros, não envia o formulário
        if (Object.keys(validationErrors).length > 0) return;

        // Marcar como enviando o formulário
        setIsSubmitting(true);

        try {
            const response = await fetch("http://localhost:5050/v1/cadastro", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    nome: formData.nome.toLowerCase(),
                    email: formData.email,
                    telefone: formData.telefone.replace(/\D/g, ""), // Remove caracteres especiais
                    senha: formData.senha,
                }),
            });
    
            const data = await response.json();
    
            if (response.ok) {
                // Exibir a mensagem de sucesso
                setIsSuccess(true);
                setIsSubmitting(false);
                
                // Esperar 3 segundos e redirecionar para a página de agendamento
                setTimeout(() => {
                    router.push("/");
                }, 3000); // Redireciona após 3 segundos
            } else {
                // Se a resposta não for ok (erro), trata os erros
                if (data.message && data.message.includes("email")) {
                    setErrors({ ...errors, email: "O email já está cadastrado." });
                }
                setIsSubmitting(false); // Garantir que o botão volte ao normal
            }
    

        } catch (error) {
            console.error("Erro ao enviar dados:", error);
            setIsSubmitting(false);
        }
    };
    

    return (
        <>
            <Navbar />
            <section id="cadastro" className="bg-cadastro bg-cover min-h-screen flex flex-col justify-center items-center">
                {/* Overlay Fundo */}
                <div className="overlay flex flex-col justify-center items-center bg-mainColor/75 absolute w-full h-full">
                    <div className="cadastro flex flex-col md:flex-row items-center justify-center gap-8 p-6 md:p-12 mt-16">
                        
                        {/* Box Esquerda */}
                        <div className="box-esquerda flex-1 text-center md:text-left px-4 md:px-[10%]">
                            <h1 className="text-xl text-center md:text-3xl text-white mb-6">
                                Com seu <span className="text-white">CADASTRO</span> você concorre a <span className="text-laranja">PRÊMIOS</span> e <span className="text-laranja">PROMOÇÕES</span> imperdíveis.
                            </h1>
                            <h1 className="text-xl text-center md:text-3xl text-white mb-6">
                                Além de receber <span className="text-laranja">AVISOS</span> sobre seus <span className="text-white">HORÁRIOS</span> e <span className="text-laranja">IMPREVISTOS</span>.
                            </h1>
                        </div>

                        {/* Box Direita - Formulário */}
                        <div className="box-direita flex-1 flex items-center justify-center p-6 rounded-xl w-full max-w-md">
                            <form onSubmit={handleSubmit} className="w-full font-archivo">
                                <h2 className="text-white text-2xl font-archivo-black mb-6 text-center">CADASTRE-SE</h2>

                                {/* Nome Completo */}
                                <div className="mb-4">
                                    <input  
                                        type="text" 
                                        name="nome" 
                                        value={formData.nome} 
                                        placeholder="Nome completo"
                                        onChange={handleChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.nome && <p className="text-red-500 text-sm">{errors.nome}</p>}
                                </div>

                                {/* E-mail */}
                                <div className="mb-4">
                                    <input 
                                        type="email" 
                                        name="email" 
                                        value={formData.email}
                                        placeholder="Email" 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.email && <p className="text-red-500 text-sm">{errors.email}</p>}
                                </div>

                                {/* Telefone */}
                                <div className="mb-4">
                                    <input 
                                        type="tel" 
                                        name="telefone" 
                                        value={formData.telefone} 
                                        placeholder="(xx) xxxxx-xxxx"
                                        onInput={handlePhoneChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.telefone && <p className="text-red-500 text-sm">{errors.telefone}</p>}
                                </div>

                                {/* Senha */}
                                <div className="mb-4">
                                    <input 
                                        type="password" 
                                        name="senha" 
                                        value={formData.senha}
                                        placeholder="Senha" 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.senha && <p className="text-red-500 text-sm">{errors.senha}</p>}
                                </div>

                                {/* Confirmar Senha */}
                                <div className="mb-4">
                                    <input 
                                        type="password" 
                                        name="confirmarSenha" 
                                        value={formData.confirmarSenha} 
                                        placeholder="Confirmar senha"
                                        onChange={handleChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.confirmarSenha && <p className="text-red-500 text-sm">{errors.confirmarSenha}</p>}
                                </div>

                                {/* Aceitar Termos */}
                                <div className="mb-4 flex items-center">
                                    <input 
                                        type="checkbox" 
                                        name="termos" 
                                        checked={formData.termos} 
                                        onChange={handleChange} 
                                        className="w-5 h-5 text-laranja bg-gray-700 border-none rounded focus:ring-2 focus:ring-laranja"
                                    />
                                    <label className="text-white ml-2">
                                        Eu concordo com os <a href="/termos" className="text-laranja underline" target="_blank">termos e condições</a>
                                    </label>
                                </div>
                                {errors.termos && <p className="text-red-500 text-sm">{errors.termos}</p>}

                                {/* Botão de Cadastro */}
                                <button 
                                    type="submit" 
                                    className="w-full bg-laranja font-archivo-black text-white py-3 rounded-full font-bold text-lg hover:bg-orange-500 transition"
                                    disabled={isSubmitting}
                                >
                                    {isSubmitting ? 'Enviando...' : 'CADASTRAR'}
                                </button>
                            </form>
                        </div>
                    </div>
                </div>

                {/* Mensagem de sucesso */}
                {isSuccess && (
                    <div className="fixed bottom-10 left-0 right-0 bg-green-500 text-white text-center py-3">
                        Cadastro realizado com sucesso! Redirecionando...
                    </div>
                )}
            </section>
        </>
    );
}
