import Hero from "~/components/landing/hero";
import ScrollTransition from "~/components/landing/scrollTransition";
import FeaturesSection from "~/components/landing/features";
import CTASection from "~/components/landing/cta";
import Footer from "~/components/landing/footer";

export function meta() {
    return [{ title: "TigerJunction" }];
}

export default function Landing() {
    return (
        <div className="relative">
            <div className="fixed inset-0 bg-gradient-to-b from-white via-gray-50 to-white -z-50" />
            <Hero />
            <ScrollTransition />
            <FeaturesSection />
            <CTASection />
            <Footer />
        </div>
    );
}
