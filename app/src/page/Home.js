import React, { useState, useEffect } from "react";
import { fetchAPI, serverAddr, availableProductsURL } from "../utils/Axios";
import Hero from "../component/Hero";
import ProductGrid from "../component/ProductGrid";

const IntroSection = () => (
  <div style={{ padding: "0 15%", color: "black", fontSize: "20px" }}>
    <h2 style={{ color: "black", fontSize: "30px" }}>Partez l'esprit tranquille.</h2>
    <p>
      Choisissez vos dates de trek. Réservez votre kit en ligne et recevez le en point
      relais de votre choix en France. Le kit sera disponible 2 jours avant votre départ
      et vous avez 2 jours pour le ramener après la fin de votre trek. ‍ Bonus ? Vous
      pouvez même renvoyer votre kit depuis votre lieu de trek.
    </p>
  </div>
);

function Home() {
  const [trekInfo, setTrekInfo] = useState({});
  const [products, setProducts] = useState(null);
  const [cart, setCart] = useState([]);

  // Asynchronous call to api to get choice's corresponding products
  useEffect(() => {
    let { category, fromDate, toDate } = trekInfo;
    if (category !== undefined && fromDate instanceof Date && toDate instanceof Date) {
      // server accepts dates on format YYYY-MM-DD
      fromDate = fromDate.toISOString().substr(0, 10);
      toDate = toDate.toISOString().substr(0, 10);
      fetchAPI(
        serverAddr +
          availableProductsURL +
          category +
          `
      &fromDate=${fromDate}&toDate=${toDate}`,
        setProducts
      );
    }
  }, [trekInfo]);

  return (
    <>
      <Hero setTrek={setTrekInfo}></Hero>
      {products === null ? (
        <IntroSection />
      ) : (
        <ProductGrid products={products} cartProducts={cart} addToCart={setCart} />
      )}
    </>
  );
}

export default Home;
