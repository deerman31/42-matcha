import type { MetaFunction } from "@remix-run/node";
import { Link } from "@remix-run/react";
import { Users, MessageCircle, Shield } from "lucide-react";
import React from "react";

export const meta: MetaFunction = () => {
  return [
    { title: "無事之名馬 New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};


export default function Index() {
  return (
    <div className="max-w-7xl mx-auto">
      {/* Hero Section */}
      <section className="text-center py-16 px-4">
        <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
          Find Your Perfect Match
        </h1>
        <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
          Matcha brings people together through intelligent matching and genuine connections.
          Start your journey to find meaningful relationships today.
        </p>
        <Link
          to="/auth/signup"
          className="inline-block bg-rose-500 text-white px-8 py-4 rounded-full text-lg font-semibold hover:bg-rose-600 transition-colors"
        >
          Sign up
        </Link>
      </section>

      {/* Features Section */}
      <section className="py-16 px-4 bg-gray-50">
        <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">
          Why Choose Matcha
        </h2>
        <div className="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto">
          <FeatureCard
            icon={<Users className="w-8 h-8 text-rose-500" />}
            title="Smart Matching"
            description="Our intelligent algorithm finds compatible matches based on your interests, preferences, and location."
          />
          <FeatureCard
            icon={<MessageCircle className="w-8 h-8 text-rose-500" />}
            title="Real-time Chat"
            description="Connect instantly with your matches through our real-time messaging system."
          />
          <FeatureCard
            icon={<Shield className="w-8 h-8 text-rose-500" />}
            title="Safe & Secure"
            description="Your privacy and security are our top priority. All profiles are verified and data is encrypted."
          />
        </div>
      </section>

      {/* How It Works Section */}
      <section className="py-16 px-4">
        <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">
          How Matcha Works
        </h2>
        <div className="grid md:grid-cols-4 gap-8 max-w-5xl mx-auto">
          <Step
            number="1"
            title="Create Profile"
            description="Sign up and create your detailed profile"
          />
          <Step
            number="2"
            title="Set Preferences"
            description="Tell us what you're looking for"
          />
          <Step
            number="3"
            title="Find Matches"
            description="Browse through compatible profiles"
          />
          <Step
            number="4"
            title="Connect"
            description="Start chatting with your matches"
          />
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 px-4 bg-rose-50 text-center">
        <h2 className="text-3xl font-bold text-gray-900 mb-6">
          Ready to Find Your Match?
        </h2>
        <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
          Join thousands of people who have already found meaningful connections on Matcha.
        </p>
        <Link
          to="/signup"
          className="inline-block bg-rose-500 text-white px-8 py-4 rounded-full text-lg font-semibold hover:bg-rose-600 transition-colors"
        >
          Start Your Journey
        </Link>
      </section>
    </div>
  );
}
interface FeatureCardProps {
  icon: React.ReactNode;
  title: string;
  description: string;
}

// Feature Card Component
const FeatureCard: React.FC<FeatureCardProps> = ({ icon, title, description }) => {
  return (
    <div className="p-6 bg-white rounded-xl shadow-sm">
      <div className="mb-4">
        {icon}
      </div>
      <h3 className="text-xl font-semibold text-gray-900 mb-2">{title}</h3>
      <p className="text-gray-600">{description}</p>
    </div>
  );
};

interface StepProps {
  number: string;
  title: string;
  description: string;
}
// Step Component
const Step: React.FC<StepProps> = ({ number, title, description }) => {
  return (
    <div className="text-center">
      <div className="w-12 h-12 bg-rose-500 text-white rounded-full flex items-center justify-center text-xl font-bold mx-auto mb-4">
        {number}
      </div>
      <h3 className="text-xl font-semibold text-gray-900 mb-2">{title}</h3>
      <p className="text-gray-600">{description}</p>
    </div>
  );
};