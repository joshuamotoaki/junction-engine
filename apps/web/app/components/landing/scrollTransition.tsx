import React, { useRef } from "react";
import { motion, useScroll, useTransform } from "framer-motion";

const ScrollTransition = () => {
    const containerRef = useRef(null);
    const { scrollYProgress } = useScroll({
        target: containerRef,
        offset: ["start end", "end start"],
    });

    return (
        <div
            ref={containerRef}
            className="relative h-[150vh] bg-gradient-to-b from-white to-gray-50"
        >
            {/* 3D rotating planes */}
            <div className="sticky top-0 h-screen overflow-hidden perspective-[1000px]">
                <motion.div
                    className="absolute inset-0 flex items-center justify-center"
                    style={{
                        rotateX: useTransform(
                            scrollYProgress,
                            [0, 0.5],
                            [0, -60]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0, 0.5],
                            [1, 0.8]
                        ),
                        y: useTransform(scrollYProgress, [0, 0.5], [0, -100]),
                    }}
                >
                    <div className="relative">
                        <motion.h2
                            className="text-[15vw] font-black text-gray-200 leading-none"
                            style={{
                                opacity: useTransform(
                                    scrollYProgress,
                                    [0, 0.3],
                                    [1, 0]
                                ),
                            }}
                        >
                            TIGER
                        </motion.h2>
                    </div>
                </motion.div>

                <motion.div
                    className="absolute inset-0 flex items-center justify-center"
                    style={{
                        rotateX: useTransform(
                            scrollYProgress,
                            [0.2, 0.7],
                            [60, 0]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0.2, 0.7],
                            [0.8, 1]
                        ),
                        y: useTransform(scrollYProgress, [0.2, 0.7], [100, 0]),
                        opacity: useTransform(
                            scrollYProgress,
                            [0.2, 0.5],
                            [0, 1]
                        ),
                    }}
                >
                    <div className="text-center">
                        <h2 className="text-[10vw] font-light text-gray-900 leading-none tracking-tighter">
                            EXPLORE
                        </h2>
                        <div className="mt-4 flex items-center justify-center gap-4">
                            <div className="w-20 h-[1px] bg-gray-400" />
                            <span className="text-xs tracking-[0.2em] text-gray-500">
                                THE POSSIBILITIES
                            </span>
                            <div className="w-20 h-[1px] bg-gray-400" />
                        </div>
                    </div>
                </motion.div>

                {/* Floating 3D elements */}
                <motion.div
                    className="absolute top-1/4 left-1/4 w-32 h-32 border border-gray-300"
                    style={{
                        rotateY: useTransform(
                            scrollYProgress,
                            [0, 1],
                            [0, 360]
                        ),
                        rotateX: useTransform(
                            scrollYProgress,
                            [0, 1],
                            [0, 360]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0, 0.5, 1],
                            [0.5, 1, 0.5]
                        ),
                    }}
                />

                <motion.div
                    className="absolute bottom-1/4 right-1/4 w-24 h-24 bg-gray-100 rounded-full"
                    style={{
                        rotateZ: useTransform(
                            scrollYProgress,
                            [0, 1],
                            [0, -360]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0, 0.5, 1],
                            [0.3, 1, 0.3]
                        ),
                        opacity: useTransform(
                            scrollYProgress,
                            [0, 0.2, 0.8, 1],
                            [0, 1, 1, 0]
                        ),
                    }}
                />
            </div>
        </div>
    );
};

export default ScrollTransition;
