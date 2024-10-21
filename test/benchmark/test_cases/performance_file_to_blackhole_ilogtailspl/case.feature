@input
Feature: performance file to blackhole iLogtail
  Performance file to blackhole iLogtail

  @e2e-performance @docker-compose
  Scenario: PerformanceFileToBlackholeiLogtailSPL
    Given {docker-compose} environment
    Given docker-compose boot type {benchmark}
    When start docker-compose {performance_file_to_blackhole_ilogtailspl}
    When start monitor {ilogtailC}
    When generate random json logs to file, speed {10}MB/s, total {1}min, to file {./a.log}
    When wait monitor until log processing finished