# hopcount
**Distance Aware Routing for Hop Efficient Systems**

* Author: Vijaya Krishna Namala
* Published In : International Journal of Technology and Applied Science (IJTAS)
* Publication Date: Feb 2021
* E-ISSN: 2230-9004
* Impact Factor : 10.31

**Abstract:**
Distributed congestion monitoring systems often consume excessive memory due to independent telemetry storage and duplication across nodes. This work proposes a memory efficient framework that consolidates telemetry handling and minimizes redundant buffering. By introducing shared aggregation and lightweight data collection, the approach significantly reduces runtime memory overhead. Experimental results across multiple cluster sizes demonstrate improved scalability and efficient resource utilization in distributed environments.

**Key Contributions**
* **Memory Efficient Monitoring Framework Design:**\
Developed a distributed monitoring architecture that minimizes memory overhead by eliminating redundant telemetry storage across nodes.
* **Consolidated Telemetry Storage Mechanism:**\
Introduced a shared aggregation approach where telemetry data is centrally maintained, avoiding duplication of buffers, logs, and intermediate data across machines.
* **Reduced Runtime Memory Overhead Strategy:**\
Designed lightweight data collection techniques that limit local buffering and reduce unnecessary memory allocation during monitoring operations.
* **Scalability Evaluation Across Cluster Sizes:**\
Conducted experiments on clusters with 3, 5, 7, 9, and 11 nodes to analyze memory consumption behavior and validate improved scalability.
**Relevance & Real World Impact**
* **Reduced Memory Consumption :**\
The proposed approach significantly lowers memory usage by avoiding duplicated telemetry storage and optimizing buffer management in distributed monitoring systems.
* **Improved Resource Utilization :**\
Efficient memory handling ensures that more system resources remain available for application workloads, improving overall system performance.
* **Enhanced Scalability in Distributed Environments :**\
Controlled memory growth with increasing cluster size enables scalable monitoring without excessive resource overhead.
* **Stable System Performance :**\
Lower memory pressure reduces risks of garbage collection delays and paging, ensuring consistent and reliable monitoring behavior.
* **Academic and Practical Contribution :**\
Provides a structured approach for designing resource efficient monitoring systems, supporting further research and real world implementation in cloud and distributed infrastructures.

**Experimental Results (Summary)**:

  | Nodes | local monitoring (MB) | Runtime Optimized (MB)| Improvement (%) |
  |-------|-----------------------| ----------------------| ----------------|
  | 3     | 920                   | 700                   | 23.91           |
  | 5     | 1100                  | 820                   | 25.45           |
  | 7     | 1280                  | 940                   | 26.56           |
  | 9     | 1460                  | 1050                  | 28.08           |
  | 11    | 1640                  | 1180                  | 28.05           |

**Citation** \
Reducing Runtime Overhead in Distributed Congestion Monitoring Systems. \
Vijaya Krishna Namala \
International Journal of Technology and Applied Science \
E-ISSN- 2230-9004  \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: linkedin.com/in/vijaya-krishna-namala-a42b2958 | **Email**: vijaya.namala@gmail.com



