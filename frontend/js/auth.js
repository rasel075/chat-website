// Handle login form submission
function handleLoginSubmit() {
	const form = document.getElementById("loginForm");
	form.addEventListener("submit", async (e) => {
		e.preventDefault();

		const email = document.getElementById("loginEmail").value;
		const password = document.getElementById("loginPassword").value;
		const errorDiv = document.getElementById("loginError");

		// Clear previous error
		errorDiv.classList.remove("show");

		try {
			// Call API
			const response = await loginUser(email, password);

			// Store token
			localStorage.setItem("token", response.token);

			// Store user info
			localStorage.setItem("userId", response.user.id);
			localStorage.setItem("username", response.user.username);

			// Redirect to chat page
			window.location.href = "chat.html";
		} catch (error) {
			// Show error message
			errorDiv.textContent = error.message;
			errorDiv.classList.add("show");
		}
	});
}

// Handle register form submission
function handleRegisterSubmit() {
	const form = document.getElementById("registerForm");
	form.addEventListener("submit", async (e) => {
		e.preventDefault();

		const username = document.getElementById("registerUsername").value;
		const email = document.getElementById("registerEmail").value;
		const password = document.getElementById("registerPassword").value;
		const confirmPassword = document.getElementById("registerConfirmPassword").value;
		const errorDiv = document.getElementById("registerError");
		const successDiv = document.getElementById("registerSuccess");

		// Clear previous messages
		errorDiv.classList.remove("show");
		successDiv.classList.remove("show");

		// Validate passwords match
		if (password !== confirmPassword) {
			errorDiv.textContent = "Passwords do not match";
			errorDiv.classList.add("show");
			return;
		}

		try {
			// Call API
			await registerUser(username, email, password);

			// Show success message
			successDiv.textContent = "Registration successful! Redirecting to login...";
			successDiv.classList.add("show");

			// Redirect to login after 2 seconds
			setTimeout(() => {
				window.location.href = "login.html";
			}, 2000);
		} catch (error) {
			// Show error message
			errorDiv.textContent = error.message;
			errorDiv.classList.add("show");
		}
	});
}

// Check if user is logged in
function checkAuth() {
	const token = localStorage.getItem("token");
	if (!token) {
		// Redirect to login if not authenticated
		window.location.href = "login.html";
		return false;
	}
	return true;
}

// Logout user
function logout() {
	localStorage.clear();
	window.location.href = "login.html";
}

// Initialize auth page
function initAuthPage() {
	const loginForm = document.getElementById("loginForm");
	const registerForm = document.getElementById("registerForm");

	if (loginForm) {
		handleLoginSubmit();
	}

	if (registerForm) {
		handleRegisterSubmit();
	}

	// Add toggle between login and register
	const toggleLinks = document.querySelectorAll(".auth-toggle");
	toggleLinks.forEach((link) => {
		link.addEventListener("click", (e) => {
			e.preventDefault();
			const loginContainer = document.getElementById("loginContainer");
			const registerContainer = document.getElementById("registerContainer");

			if (loginContainer.style.display === "none") {
				loginContainer.style.display = "block";
				registerContainer.style.display = "none";
			} else {
				loginContainer.style.display = "none";
				registerContainer.style.display = "block";
			}
		});
	});
}

// Call init when page loads
if (document.readyState === "loading") {
	document.addEventListener("DOMContentLoaded", initAuthPage);
} else {
	initAuthPage();
}